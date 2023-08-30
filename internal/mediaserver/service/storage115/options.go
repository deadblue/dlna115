package storage115

import (
	"bufio"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/deadblue/dlna115/internal/util"
	"github.com/deadblue/elevengo"
)

const (
	CredentialSourceFile = "file"
	CredentialSourceUrl  = "url"

	FolderTypeDir   = "dir"
	FolderTypeStar  = "star"
	FolderTypeLabel = "label"
)

var (
	ErrUnsupportedCredentialSource = errors.New("unsupported credential source")
)

type CredentialSourceOption struct {
	Type   string `yaml:"type"`
	Source string `yaml:"source"`
}

type TopFolderOption struct {
	Type   string `yaml:"type"`
	Name   string `yaml:"name"`
	Target string `yaml:"target"`
}

type Options struct {
	// Credential source
	CredentialSource CredentialSourceOption `yaml:"credential-source"`
	// Top folder config
	TopFolders []TopFolderOption `yaml:"top-folders"`
}

func (s *Service) ApplyOptions() (err error) {
	if err = s.loadCredential(); err != nil {
		return
	}
	if err = s.initTopFolders(); err != nil {
		return
	}
	return
}

func (s *Service) loadCredential() (err error) {
	src := &s.opts.CredentialSource
	cred := &elevengo.Credential{}

	// Open stream of source
	var r io.ReadCloser
	switch src.Type {
	case CredentialSourceFile:
		r, err = os.Open(src.Source)
	case CredentialSourceUrl:
		var resp *http.Response
		resp, err = http.Get(src.Source)
		r = resp.Body
	default:
		err = ErrUnsupportedCredentialSource
	}
	if err != nil {
		return
	}
	defer r.Close()

	// The credential content shoule be in this form:
	//		UID=UID_from_Cookie
	//		CID=CID_from_Cookie
	//		SEID=SEID_from_Cookie

	// Read file line by line
	for s := bufio.NewScanner(r); s.Scan(); {
		line := s.Text()
		sepIndex := strings.IndexRune(line, '=')
		name := strings.TrimSpace(line[:sepIndex])
		value := strings.TrimSpace(line[sepIndex+1:])
		switch name {
		case "UID":
			cred.UID = value
		case "CID":
			cred.CID = value
		case "SEID":
			cred.SEID = value
		}
	}
	return s.ea.CredentialImport(cred)
}

func (s *Service) initTopFolders() (err error) {
	// Get all labels
	labelMap := make(map[string]string)
	it, err := s.ea.LabelIterate()
	if err != nil {
		return
	}
	for ; err == nil; err = it.Next() {
		label := &elevengo.Label{}
		if it.Get(label) == nil {
			labelMap[label.Name] = label.Id
		}
	}
	err = nil

	for _, tfo := range s.opts.TopFolders {
		tf := &Folder{
			Type: tfo.Type,
		}
		switch tfo.Type {
		case FolderTypeStar:
			tf.Type = FolderTypeStar
			tf.Name = util.DefaultString(tfo.Name, "Favorites")
		case FolderTypeLabel:
			if labelId, ok := labelMap[tfo.Target]; ok {
				tf.SourceId = labelId
				tf.Name = util.DefaultString(tfo.Name, tfo.Target)
			} else {
				tf = nil
			}
		case FolderTypeDir:
			if dirId, err := s.ea.DirGetId(tfo.Target); err == nil {
				tf.SourceId = dirId
				tf.Name = util.DefaultStringFunc(
					tfo.Name, func() string {
						return filepath.Base(tfo.Target)
					},
				)
			} else {
				tf = nil
			}
		default:
			log.Printf("Unsupported folder type: %s", tfo.Type)
			tf = nil
		}
		if tf != nil {
			s.tfs = append(s.tfs, tf)
		}
	}
	return
}
