package main

// Webp - Image format providing lossless and lossy compression for web images
// Homepage: https://developers.google.com/speed/webp/

import (
	"fmt"
	
	"os/exec"
)

func installWebp() {
	// Método 1: Descargar y extraer .tar.gz
	webp_tar_url := "https://storage.googleapis.com/downloads.webmproject.org/releases/webp/libwebp-1.4.0.tar.gz"
	webp_cmd_tar := exec.Command("curl", "-L", webp_tar_url, "-o", "package.tar.gz")
	err := webp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	webp_zip_url := "https://storage.googleapis.com/downloads.webmproject.org/releases/webp/libwebp-1.4.0.zip"
	webp_cmd_zip := exec.Command("curl", "-L", webp_zip_url, "-o", "package.zip")
	err = webp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	webp_bin_url := "https://storage.googleapis.com/downloads.webmproject.org/releases/webp/libwebp-1.4.0.bin"
	webp_cmd_bin := exec.Command("curl", "-L", webp_bin_url, "-o", "binary.bin")
	err = webp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	webp_src_url := "https://storage.googleapis.com/downloads.webmproject.org/releases/webp/libwebp-1.4.0.src.tar.gz"
	webp_cmd_src := exec.Command("curl", "-L", webp_src_url, "-o", "source.tar.gz")
	err = webp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	webp_cmd_direct := exec.Command("./binary")
	err = webp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: giflib")
exec.Command("latte", "install", "giflib")
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: libtiff")
exec.Command("latte", "install", "libtiff")
}
