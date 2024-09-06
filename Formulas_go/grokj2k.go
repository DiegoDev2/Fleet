package main

// Grokj2k - JPEG 2000 Library
// Homepage: https://github.com/GrokImageCompression/grok

import (
	"fmt"
	
	"os/exec"
)

func installGrokj2k() {
	// Método 1: Descargar y extraer .tar.gz
	grokj2k_tar_url := "https://github.com/GrokImageCompression/grok.git"
	grokj2k_cmd_tar := exec.Command("curl", "-L", grokj2k_tar_url, "-o", "package.tar.gz")
	err := grokj2k_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	grokj2k_zip_url := "https://github.com/GrokImageCompression/grok.git"
	grokj2k_cmd_zip := exec.Command("curl", "-L", grokj2k_zip_url, "-o", "package.zip")
	err = grokj2k_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	grokj2k_bin_url := "https://github.com/GrokImageCompression/grok.git"
	grokj2k_cmd_bin := exec.Command("curl", "-L", grokj2k_bin_url, "-o", "binary.bin")
	err = grokj2k_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	grokj2k_src_url := "https://github.com/GrokImageCompression/grok.git"
	grokj2k_cmd_src := exec.Command("curl", "-L", grokj2k_src_url, "-o", "source.tar.gz")
	err = grokj2k_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	grokj2k_cmd_direct := exec.Command("./binary")
	err = grokj2k_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: doxygen")
exec.Command("latte", "install", "doxygen")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: exiftool")
exec.Command("latte", "install", "exiftool")
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: libtiff")
exec.Command("latte", "install", "libtiff")
	fmt.Println("Instalando dependencia: little-cms2")
exec.Command("latte", "install", "little-cms2")
	fmt.Println("Instalando dependencia: llvm")
exec.Command("latte", "install", "llvm")
	fmt.Println("Instalando dependencia: xz")
exec.Command("latte", "install", "xz")
	fmt.Println("Instalando dependencia: zstd")
exec.Command("latte", "install", "zstd")
}
