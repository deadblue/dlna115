package storage115

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/deadblue/dlna115/pkg/credential"
	"github.com/deadblue/dlna115/pkg/util"
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
	Secret string `yaml:"secret,omitempty"`
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
	TopFolders []TopFolderOption `yaml:"top-folders,omitempty"`
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

	// Read credential data
	credData, err := io.ReadAll(r)
	if err != nil {
		return
	}
	// Decode credential
	cred := &elevengo.Credential{}
	if err = credential.Decode(credData, src.Secret, cred); err == nil {
		err = s.ea.CredentialImport(cred)
	}
	return
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

	if s.opts.TopFolders != nil {
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
				log.Printf("Add top folder: %s", tf.Name)
				s.tfs = append(s.tfs, tf)
			}
		}
	}
	// Add default top folders when user not set
	if len(s.tfs) == 0 {
		s.tfs = []*Folder{
			{
				Type: FolderTypeStar,
				Name: "Favorites",
			},
			{
				Type:     FolderTypeDir,
				Name:     "All Files",
				SourceId: "0",
			},
		}
	}
	return
}
