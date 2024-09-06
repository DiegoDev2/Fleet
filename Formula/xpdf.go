package main

// Xpdf - PDF viewer
// Homepage: https://www.xpdfreader.com/

import (
	"fmt"
	
	"os/exec"
)

func installXpdf() {
	// Método 1: Descargar y extraer .tar.gz
	xpdf_tar_url := "https://dl.xpdfreader.com/xpdf-4.05.tar.gz"
	xpdf_cmd_tar := exec.Command("curl", "-L", xpdf_tar_url, "-o", "package.tar.gz")
	err := xpdf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xpdf_zip_url := "https://dl.xpdfreader.com/xpdf-4.05.zip"
	xpdf_cmd_zip := exec.Command("curl", "-L", xpdf_zip_url, "-o", "package.zip")
	err = xpdf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xpdf_bin_url := "https://dl.xpdfreader.com/xpdf-4.05.bin"
	xpdf_cmd_bin := exec.Command("curl", "-L", xpdf_bin_url, "-o", "binary.bin")
	err = xpdf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xpdf_src_url := "https://dl.xpdfreader.com/xpdf-4.05.src.tar.gz"
	xpdf_cmd_src := exec.Command("curl", "-L", xpdf_src_url, "-o", "source.tar.gz")
	err = xpdf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xpdf_cmd_direct := exec.Command("./binary")
	err = xpdf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: fontconfig")
	exec.Command("latte", "install", "fontconfig").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: qt@5")
	exec.Command("latte", "install", "qt@5").Run()
}
