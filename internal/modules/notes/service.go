package notes

type NoteService interface {
	FindAllNotes() ([]Note, error)
	FindByIDNote(id string) (*Note, error)
	CreateNote(note *Note) error
	UpdateNote(note *Note) error
	DeleteNote(id string) error
}

// Struct
type noteService struct {
	noteRepository NoteRepository
}

// Tạo mới service
func NewNoteService(noteRepository NoteRepository) NoteService {
	return &noteService{noteRepository: noteRepository}
}

// Tìm tất cả notes
func (s *noteService) FindAllNotes() ([]Note, error) {
	return s.noteRepository.FindAllNotes()
}

// Tìm note theo ID
func (s *noteService) FindByIDNote(id string) (*Note, error) {
	return s.noteRepository.FindByIDNote(id)
}

// Tạo mới note
func (s *noteService) CreateNote(note *Note) error {
	return s.noteRepository.CreateNote(note)
}

// Cập nhật note
func (s *noteService) UpdateNote(note *Note) error {
	return s.noteRepository.UpdateNote(note)
}

// Xóa note
func (s *noteService) DeleteNote(id string) error {
	return s.noteRepository.DeleteNote(id)
}
