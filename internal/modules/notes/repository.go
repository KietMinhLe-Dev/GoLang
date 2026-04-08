package notes

import "gorm.io/gorm"

type NoteRepository interface {
	FindAllNotes() ([]Note, error)
	FindByIDNote(id string) (*Note, error)
	CreateNote(note *Note) error
	UpdateNote(note *Note) error
	DeleteNote(id string) error
}

type noteRepository struct {
	db *gorm.DB
}

func NewNoteRepository(db *gorm.DB) NoteRepository {
	return &noteRepository{db: db}
}

func (r *noteRepository) FindAllNotes() ([]Note, error) {
	var notes []Note
	err := r.db.Find(&notes).Error
	return notes, err
}

func (r *noteRepository) FindByIDNote(id string) (*Note, error) {
	var note Note
	err := r.db.Where("id = ?", id).First(&note).Error
	return &note, err
}

func (r *noteRepository) CreateNote(note *Note) error {
	return r.db.Create(note).Error
}

func (r *noteRepository) UpdateNote(note *Note) error {
	return r.db.Save(note).Error
}

func (r *noteRepository) DeleteNote(id string) error {
	return r.db.Where("id = ?", id).Delete(&Note{}).Error
}
