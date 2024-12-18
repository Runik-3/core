export interface DictFile {
  Name: string;
  Display: string;
  Extension: string;
  Size: number;
  // ISO string
  Modified: string;
}

export interface Definition {
  Word: string
  Definition: string
}

// modifies the type signature of Definition so we can edit in-place
export interface EditableDefinition extends Definition {
  initWord: string;
  initDefinition: string;
}

export interface Dict {
  Name: string
  Lexicon: Definition[]
}
