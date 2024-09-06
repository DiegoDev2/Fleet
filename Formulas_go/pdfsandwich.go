package main

// Pdfsandwich - Generate sandwich OCR PDFs from scanned file
// Homepage: http://www.tobias-elze.de/pdfsandwich/

import (
	"fmt"
	
	"os/exec"
)

func installPdfsandwich() {
	// Método 1: Descargar y extraer .tar.gz
	pdfsandwich_tar_url := "https://downloads.sourceforge.net/project/pdfsandwich/pdfsandwich%200.1.7/pdfsandwich-0.1.7.tar.bz2"
	pdfsandwich_cmd_tar := exec.Command("curl", "-L", pdfsandwich_tar_url, "-o", "package.tar.gz")
	err := pdfsandwich_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pdfsandwich_zip_url := "https://downloads.sourceforge.net/project/pdfsandwich/pdfsandwich%200.1.7/pdfsandwich-0.1.7.tar.bz2"
	pdfsandwich_cmd_zip := exec.Command("curl", "-L", pdfsandwich_zip_url, "-o", "package.zip")
	err = pdfsandwich_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pdfsandwich_bin_url := "https://downloads.sourceforge.net/project/pdfsandwich/pdfsandwich%200.1.7/pdfsandwich-0.1.7.tar.bz2"
	pdfsandwich_cmd_bin := exec.Command("curl", "-L", pdfsandwich_bin_url, "-o", "binary.bin")
	err = pdfsandwich_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pdfsandwich_src_url := "https://downloads.sourceforge.net/project/pdfsandwich/pdfsandwich%200.1.7/pdfsandwich-0.1.7.tar.bz2"
	pdfsandwich_cmd_src := exec.Command("curl", "-L", pdfsandwich_src_url, "-o", "source.tar.gz")
	err = pdfsandwich_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pdfsandwich_cmd_direct := exec.Command("./binary")
	err = pdfsandwich_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gawk")
exec.Command("latte", "install", "gawk")
	fmt.Println("Instalando dependencia: ocaml")
exec.Command("latte", "install", "ocaml")
	fmt.Println("Instalando dependencia: exact-image")
exec.Command("latte", "install", "exact-image")
	fmt.Println("Instalando dependencia: ghostscript")
exec.Command("latte", "install", "ghostscript")
	fmt.Println("Instalando dependencia: imagemagick")
exec.Command("latte", "install", "imagemagick")
	fmt.Println("Instalando dependencia: poppler")
exec.Command("latte", "install", "poppler")
	fmt.Println("Instalando dependencia: tesseract")
exec.Command("latte", "install", "tesseract")
	fmt.Println("Instalando dependencia: unpaper")
exec.Command("latte", "install", "unpaper")
}
