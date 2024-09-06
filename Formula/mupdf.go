package main

// Mupdf - Lightweight PDF and XPS viewer
// Homepage: https://mupdf.com/

import (
	"fmt"
	
	"os/exec"
)

func installMupdf() {
	// Método 1: Descargar y extraer .tar.gz
	mupdf_tar_url := "https://mupdf.com/downloads/archive/mupdf-1.23.11-source.tar.gz"
	mupdf_cmd_tar := exec.Command("curl", "-L", mupdf_tar_url, "-o", "package.tar.gz")
	err := mupdf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mupdf_zip_url := "https://mupdf.com/downloads/archive/mupdf-1.23.11-source.zip"
	mupdf_cmd_zip := exec.Command("curl", "-L", mupdf_zip_url, "-o", "package.zip")
	err = mupdf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mupdf_bin_url := "https://mupdf.com/downloads/archive/mupdf-1.23.11-source.bin"
	mupdf_cmd_bin := exec.Command("curl", "-L", mupdf_bin_url, "-o", "binary.bin")
	err = mupdf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mupdf_src_url := "https://mupdf.com/downloads/archive/mupdf-1.23.11-source.src.tar.gz"
	mupdf_cmd_src := exec.Command("curl", "-L", mupdf_src_url, "-o", "source.tar.gz")
	err = mupdf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mupdf_cmd_direct := exec.Command("./binary")
	err = mupdf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: harfbuzz")
	exec.Command("latte", "install", "harfbuzz").Run()
	fmt.Println("Instalando dependencia: jbig2dec")
	exec.Command("latte", "install", "jbig2dec").Run()
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: mujs")
	exec.Command("latte", "install", "mujs").Run()
	fmt.Println("Instalando dependencia: openjpeg")
	exec.Command("latte", "install", "openjpeg").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: freeglut")
	exec.Command("latte", "install", "freeglut").Run()
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
	fmt.Println("Instalando dependencia: libxext")
	exec.Command("latte", "install", "libxext").Run()
	fmt.Println("Instalando dependencia: mesa")
	exec.Command("latte", "install", "mesa").Run()
}
