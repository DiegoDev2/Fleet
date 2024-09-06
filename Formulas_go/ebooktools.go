package main

// EbookTools - Access and convert several ebook formats
// Homepage: https://sourceforge.net/projects/ebook-tools/

import (
	"fmt"
	
	"os/exec"
)

func installEbookTools() {
	// Método 1: Descargar y extraer .tar.gz
	ebooktools_tar_url := "https://downloads.sourceforge.net/project/ebook-tools/ebook-tools/0.2.2/ebook-tools-0.2.2.tar.gz"
	ebooktools_cmd_tar := exec.Command("curl", "-L", ebooktools_tar_url, "-o", "package.tar.gz")
	err := ebooktools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ebooktools_zip_url := "https://downloads.sourceforge.net/project/ebook-tools/ebook-tools/0.2.2/ebook-tools-0.2.2.zip"
	ebooktools_cmd_zip := exec.Command("curl", "-L", ebooktools_zip_url, "-o", "package.zip")
	err = ebooktools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ebooktools_bin_url := "https://downloads.sourceforge.net/project/ebook-tools/ebook-tools/0.2.2/ebook-tools-0.2.2.bin"
	ebooktools_cmd_bin := exec.Command("curl", "-L", ebooktools_bin_url, "-o", "binary.bin")
	err = ebooktools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ebooktools_src_url := "https://downloads.sourceforge.net/project/ebook-tools/ebook-tools/0.2.2/ebook-tools-0.2.2.src.tar.gz"
	ebooktools_cmd_src := exec.Command("curl", "-L", ebooktools_src_url, "-o", "source.tar.gz")
	err = ebooktools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ebooktools_cmd_direct := exec.Command("./binary")
	err = ebooktools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: libzip")
exec.Command("latte", "install", "libzip")
}
