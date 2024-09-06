package main

// Pstoedit - Convert PostScript and PDF files to editable vector graphics
// Homepage: http://www.pstoedit.net/

import (
	"fmt"
	
	"os/exec"
)

func installPstoedit() {
	// Método 1: Descargar y extraer .tar.gz
	pstoedit_tar_url := "https://downloads.sourceforge.net/project/pstoedit/pstoedit/3.78/pstoedit-3.78.tar.gz"
	pstoedit_cmd_tar := exec.Command("curl", "-L", pstoedit_tar_url, "-o", "package.tar.gz")
	err := pstoedit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pstoedit_zip_url := "https://downloads.sourceforge.net/project/pstoedit/pstoedit/3.78/pstoedit-3.78.zip"
	pstoedit_cmd_zip := exec.Command("curl", "-L", pstoedit_zip_url, "-o", "package.zip")
	err = pstoedit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pstoedit_bin_url := "https://downloads.sourceforge.net/project/pstoedit/pstoedit/3.78/pstoedit-3.78.bin"
	pstoedit_cmd_bin := exec.Command("curl", "-L", pstoedit_bin_url, "-o", "binary.bin")
	err = pstoedit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pstoedit_src_url := "https://downloads.sourceforge.net/project/pstoedit/pstoedit/3.78/pstoedit-3.78.src.tar.gz"
	pstoedit_cmd_src := exec.Command("curl", "-L", pstoedit_src_url, "-o", "source.tar.gz")
	err = pstoedit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pstoedit_cmd_direct := exec.Command("./binary")
	err = pstoedit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: ghostscript")
exec.Command("latte", "install", "ghostscript")
	fmt.Println("Instalando dependencia: imagemagick")
exec.Command("latte", "install", "imagemagick")
	fmt.Println("Instalando dependencia: plotutils")
exec.Command("latte", "install", "plotutils")
}
