package convert

import (
	"fmt"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"io"
	"log"
	"os"
)

func ReadPdf(pdfFilePath string, outFileName string) {
	// 打开PDF文件
	ctx, err := api.ReadContextFile(pdfFilePath)
	if err != nil {
		log.Fatalf("Error reading PDF file: %v", err)
	}

	// 提取文本内容
	var extractedText []byte
	for page := 1; page <= ctx.PageCount; page++ {
		fmt.Println(page)
		reader, err := api.ExtractPage(ctx, page)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(reader)
		newStr, _ := io.ReadAll(reader)
		extractedText = append(extractedText, newStr...)
	}

	err = os.WriteFile(outFileName, extractedText, 0644)

	// 输出提取的文本内容
	fmt.Println(extractedText)

	if err != nil {
		log.Fatalf("Error writing Markdown file: %v", err)
	}

	fmt.Println("Markdown file created successfully!")
}
