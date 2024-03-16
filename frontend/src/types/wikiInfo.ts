interface Lang {
  lang: string
  langName: string
  autonym: string
  url: string
}

export interface WikiInfo {
  SiteName: string
  MainPage: string
  Lang: string
  Logo: string
  Pages: number
  Articles: number
  Languages: Lang[]
}
