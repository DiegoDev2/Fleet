package main

// Pdf2djvu - Create DjVu files from PDF files
// Homepage: https://jwilk.net/software/pdf2djvu

import (
	"fmt"
	
	"os/exec"
)

func installPdf2djvu() {
	// Método 1: Descargar y extraer .tar.gz
	pdf2djvu_tar_url := "https://github.com/jwilk/pdf2djvu/releases/download/0.9.19/pdf2djvu-0.9.19.tar.xz"
	pdf2djvu_cmd_tar := exec.Command("curl", "-L", pdf2djvu_tar_url, "-o", "package.tar.gz")
	err := pdf2djvu_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pdf2djvu_zip_url := "https://github.com/jwilk/pdf2djvu/releases/download/0.9.19/pdf2djvu-0.9.19.tar.xz"
	pdf2djvu_cmd_zip := exec.Command("curl", "-L", pdf2djvu_zip_url, "-o", "package.zip")
	err = pdf2djvu_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pdf2djvu_bin_url := "https://github.com/jwilk/pdf2djvu/releases/download/0.9.19/pdf2djvu-0.9.19.tar.xz"
	pdf2djvu_cmd_bin := exec.Command("curl", "-L", pdf2djvu_bin_url, "-o", "binary.bin")
	err = pdf2djvu_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pdf2djvu_src_url := "https://github.com/jwilk/pdf2djvu/releases/download/0.9.19/pdf2djvu-0.9.19.tar.xz"
	pdf2djvu_cmd_src := exec.Command("curl", "-L", pdf2djvu_src_url, "-o", "source.tar.gz")
	err = pdf2djvu_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pdf2djvu_cmd_direct := exec.Command("./binary")
	err = pdf2djvu_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: djvulibre")
	exec.Command("latte", "install", "djvulibre").Run()
	fmt.Println("Instalando dependencia: exiv2")
	exec.Command("latte", "install", "exiv2").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: poppler")
	exec.Command("latte", "install", "poppler").Run()
}
