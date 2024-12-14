export interface DictFile {
  Name: string;
  Display: string;
  Extension: string;
  Size: number;
  // ISO string
  Modified: string;
}

export interface Dict {
  Name: string
  Lexicon: {
    Word: string
    Definition: string
  }
}
