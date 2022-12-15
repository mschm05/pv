/////////////////////////////////////////
// crypto.go - Password termui encryption
// Mike Schilli, 2022 (m@perlmeister.com)
/////////////////////////////////////////
package main

import (
  "bytes"
  "filippo.io/age"
  "filippo.io/age/armor"
  "io"
  "os"
)

const secFile string = "test.age"

func writeEnc(txt string, pass string) error {
  recipient, err := age.NewScryptRecipient(pass)
  if err != nil {
    return err
  }

  out, err := os.OpenFile(secFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
  if err != nil {
    return err
  }
  defer out.Close()

  armorWriter := armor.NewWriter(out)
  defer armorWriter.Close()

  w, err := age.Encrypt(armorWriter, recipient)
  if err != nil {
    return err
  }
  defer w.Close()

  if _, err := io.WriteString(w, txt); err != nil {
    return err
  }

  return nil
}

func readEnc(pass string) (string, error) {
  identity, err := age.NewScryptIdentity(pass)
  if err != nil {
    return "", err
  }

  out := &bytes.Buffer{}

  in, err := os.Open(secFile)
  if err != nil {
    return "", err
  }
  defer in.Close()

  armorReader := armor.NewReader(in)

  r, err := age.Decrypt(armorReader, identity)
  if err != nil {
    return "", err
  }
  if _, err := io.Copy(out, r); err != nil {
    return "", err
  }

  return out.String(), nil
}
