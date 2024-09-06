package main

// Ghostscript - Interpreter for PostScript and PDF
// Homepage: https://www.ghostscript.com/

import (
	"fmt"
	
	"os/exec"
)

func installGhostscript() {
	// Método 1: Descargar y extraer .tar.gz
	ghostscript_tar_url := "https://github.com/ArtifexSoftware/ghostpdl-downloads/releases/download/gs10031/ghostpdl-10.03.1.tar.xz"
	ghostscript_cmd_tar := exec.Command("curl", "-L", ghostscript_tar_url, "-o", "package.tar.gz")
	err := ghostscript_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ghostscript_zip_url := "https://github.com/ArtifexSoftware/ghostpdl-downloads/releases/download/gs10031/ghostpdl-10.03.1.tar.xz"
	ghostscript_cmd_zip := exec.Command("curl", "-L", ghostscript_zip_url, "-o", "package.zip")
	err = ghostscript_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ghostscript_bin_url := "https://github.com/ArtifexSoftware/ghostpdl-downloads/releases/download/gs10031/ghostpdl-10.03.1.tar.xz"
	ghostscript_cmd_bin := exec.Command("curl", "-L", ghostscript_bin_url, "-o", "binary.bin")
	err = ghostscript_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ghostscript_src_url := "https://github.com/ArtifexSoftware/ghostpdl-downloads/releases/download/gs10031/ghostpdl-10.03.1.tar.xz"
	ghostscript_cmd_src := exec.Command("curl", "-L", ghostscript_src_url, "-o", "source.tar.gz")
	err = ghostscript_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ghostscript_cmd_direct := exec.Command("./binary")
	err = ghostscript_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: fontconfig")
	exec.Command("latte", "install", "fontconfig").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: jbig2dec")
	exec.Command("latte", "install", "jbig2dec").Run()
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: libidn")
	exec.Command("latte", "install", "libidn").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: libtiff")
	exec.Command("latte", "install", "libtiff").Run()
	fmt.Println("Instalando dependencia: little-cms2")
	exec.Command("latte", "install", "little-cms2").Run()
	fmt.Println("Instalando dependencia: openjpeg")
	exec.Command("latte", "install", "openjpeg").Run()
}
