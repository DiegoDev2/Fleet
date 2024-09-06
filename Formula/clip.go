package main

// Clip - Create high-quality charts from the command-line
// Homepage: https://github.com/asmuth/clip

import (
	"fmt"
	
	"os/exec"
)

func installClip() {
	// Método 1: Descargar y extraer .tar.gz
	clip_tar_url := "https://github.com/asmuth/clip/archive/refs/tags/v0.7.tar.gz"
	clip_cmd_tar := exec.Command("curl", "-L", clip_tar_url, "-o", "package.tar.gz")
	err := clip_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	clip_zip_url := "https://github.com/asmuth/clip/archive/refs/tags/v0.7.zip"
	clip_cmd_zip := exec.Command("curl", "-L", clip_zip_url, "-o", "package.zip")
	err = clip_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	clip_bin_url := "https://github.com/asmuth/clip/archive/refs/tags/v0.7.bin"
	clip_cmd_bin := exec.Command("curl", "-L", clip_bin_url, "-o", "binary.bin")
	err = clip_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	clip_src_url := "https://github.com/asmuth/clip/archive/refs/tags/v0.7.src.tar.gz"
	clip_cmd_src := exec.Command("curl", "-L", clip_src_url, "-o", "source.tar.gz")
	err = clip_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	clip_cmd_direct := exec.Command("./binary")
	err = clip_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: cairo")
	exec.Command("latte", "install", "cairo").Run()
	fmt.Println("Instalando dependencia: fmt")
	exec.Command("latte", "install", "fmt").Run()
	fmt.Println("Instalando dependencia: fontconfig")
	exec.Command("latte", "install", "fontconfig").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: fribidi")
	exec.Command("latte", "install", "fribidi").Run()
	fmt.Println("Instalando dependencia: harfbuzz")
	exec.Command("latte", "install", "harfbuzz").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
}
