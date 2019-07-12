package input_test

import (
	"fmt"
	"os"
	"path/filepath"

	"../input"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/afero"
)

var _ = Describe("Reader", func() {
	var (
		err    error
		fs     afero.Fs
		reader *input.ReaderImpl
	)

	BeforeEach(func() {
		fs = afero.NewMemMapFs()
	})

	Describe("NewReader", func() {
		Context("With a filesystem", func() {
			JustBeforeEach(func() {
				reader = input.NewReader(fs)
			})

			It("can create a reader", func() {
				Expect(reader).ToNot(BeNil())
			})
		})
	})

	Describe("FindFiles", func() {
		var (
			root  string
			files []string
		)

		BeforeEach(func() {
			root = "/the/file/path"
			reader = input.NewReader(fs)
		})

		JustBeforeEach(func() {
			files, err = reader.FindFiles(root)
		})

		Context("with an empty filesystem", func() {
			It("does error", func() {
				Expect(err).ToNot(BeNil())
			})
		})

		Context("with the root folder", func() {
			BeforeEach(func() {
				if err = fs.Mkdir(root, os.ModePerm); err != nil {
					Fail(err.Error())
				}
			})

			Context("with no files", func() {
				It("does not error", func() {
					Expect(err).To(BeNil())
				})
				It("finds no files", func() {
					Expect(files).To(BeEmpty())
				})
			})

			Context("with unmatched files", func() {
				BeforeEach(func() {
					for i := 0; i < 5; i++ {
						if err = afero.WriteFile(fs, filepath.Join(root, fmt.Sprintf("file%d.wtf", i)), []byte("unmatched test file"), os.ModePerm); err != nil {
							Fail(err.Error())
						}
					}
				})

				It("does not error", func() {
					Expect(err).To(BeNil())
				})
				It("finds no files", func() {
					Expect(files).To(BeEmpty())
				})
			})

			Context("with a single matched file", func() {
				BeforeEach(func() {
					for i := 0; i < 5; i++ {
						if err = afero.WriteFile(fs, filepath.Join(root, fmt.Sprintf("%d.no", i)), []byte("unmatched test file"), os.ModePerm); err != nil {
							Fail(err.Error())
						}
					}

					if err = afero.WriteFile(fs, filepath.Join(root, "yes.yaml"), []byte("this should match even though its not valid yaml"), os.ModePerm); err != nil {
						Fail(err.Error())
					}
				})

				It("does not error", func() {
					Expect(err).To(BeNil())
				})
				It("finds one file", func() {
					Expect(files).To(HaveLen(1))
				})
				It("finds the right file", func() {
					Expect(files[0]).To(Equal(filepath.Join(root, "yes.yaml")))
				})
			})

			Context("with multiple matched files", func() {
				BeforeEach(func() {
					for i := 0; i < 5; i++ {
						if err = afero.WriteFile(fs, filepath.Join(root, fmt.Sprintf("%d.no", i)), []byte("unmatched test file"), os.ModePerm); err != nil {
							Fail(err.Error())
						}
					}

					for i := 0; i < 2; i++ {
						if err = afero.WriteFile(fs, filepath.Join(root, fmt.Sprintf("%d.yaml", i)), []byte("this should match even though its not valid yaml"), os.ModePerm); err != nil {
							Fail(err.Error())
						}
					}
					for i := 2; i < 4; i++ {
						if err = afero.WriteFile(fs, filepath.Join(root, fmt.Sprintf("%d.yml", i)), []byte("this should match even though its not valid yaml"), os.ModePerm); err != nil {
							Fail(err.Error())
						}
					}
					for i := 4; i < 6; i++ {
						if err = afero.WriteFile(fs, filepath.Join(root, fmt.Sprintf("%d.comp", i)), []byte("this should match even though its not valid yaml"), os.ModePerm); err != nil {
							Fail(err.Error())
						}
					}
				})

				It("does not error", func() {
					Expect(err).To(BeNil())
				})
				It("finds one file", func() {
					Expect(files).To(HaveLen(6))
				})
				It("finds the right files", func() {
					Expect(files).To(ConsistOf([]string{
						filepath.Join(root, "0.yaml"),
						filepath.Join(root, "1.yaml"),
						filepath.Join(root, "2.yml"),
						filepath.Join(root, "3.yml"),
						filepath.Join(root, "4.comp"),
						filepath.Join(root, "5.comp"),
					}))
				})
			})
		})
	})

	Describe("ReadAll", func() {
		var (
			root  string
			bytes []byte
		)

		BeforeEach(func() {
			root = "/the/file/path"
			reader = input.NewReader(fs)
		})

		JustBeforeEach(func() {
			bytes, err = reader.ReadAll(root)
		})

		Context("with an empty filesystem", func() {
			It("does error", func() {
				Expect(err).ToNot(BeNil())
			})
		})

		Context("with the root folder", func() {
			BeforeEach(func() {
				if err = fs.Mkdir(root, os.ModePerm); err != nil {
					Fail(err.Error())
				}
			})

			Context("with a single matched file ending without a newline", func() {
				var (
					fileOne string
				)

				BeforeEach(func() {
					fileOne = "the first file content"

					if err = afero.WriteFile(fs, filepath.Join(root, "one.yaml"), []byte(fileOne), os.ModePerm); err != nil {
						Fail(err.Error())
					}
				})

				It("does not error", func() {
					Expect(err).To(BeNil())
				})
				It("reads the bytes from the file", func() {
					Expect(bytes).To(Equal([]byte(fileOne + "\n")))
				})
			})

			Context("with a single matched file ending with a newline", func() {
				var (
					fileOne string
				)

				BeforeEach(func() {
					fileOne = "the first file content\n"

					if err = afero.WriteFile(fs, filepath.Join(root, "one.yaml"), []byte(fileOne), os.ModePerm); err != nil {
						Fail(err.Error())
					}
				})

				It("does not error", func() {
					Expect(err).To(BeNil())
				})
				It("reads the bytes from the file", func() {
					Expect(bytes).To(Equal([]byte(fileOne)))
				})
			})

			Context("with multiple matched files", func() {
				var (
					fileOne   string
					fileTwo   string
					fileThree string
				)

				BeforeEach(func() {
					fileOne = `the first file content`
					fileTwo = `the second file content`
					fileThree = `
this:
  is:
    yamly: for sure
`

					if err = afero.WriteFile(fs, filepath.Join(root, "one.yaml"), []byte(fileOne), os.ModePerm); err != nil {
						Fail(err.Error())
					}
					if err = afero.WriteFile(fs, filepath.Join(root, "nested", "two.yml"), []byte(fileTwo), os.ModePerm); err != nil {
						Fail(err.Error())
					}
					if err = afero.WriteFile(fs, filepath.Join(root, "nested", "deep", "three.comp"), []byte(fileThree), os.ModePerm); err != nil {
						Fail(err.Error())
					}
				})

				It("does not error", func() {
					Expect(err).To(BeNil())
				})
				It("reads bytes from all files", func() {
					Expect(bytes).To(ContainSubstring(fileOne))
					Expect(bytes).To(ContainSubstring(fileTwo))
					Expect(bytes).To(ContainSubstring(fileThree))
				})
				It("does not read random bytes", func() {
					Expect(bytes).ToNot(ContainSubstring("42"))
				})
			})
		})
	})
})
