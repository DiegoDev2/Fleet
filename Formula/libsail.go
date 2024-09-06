package main

// Libsail - Missing small and fast image decoding library for humans (not for machines)
// Homepage: https://github.com/HappySeaFox/sail

import (
	"fmt"
	
	"os/exec"
)

func installLibsail() {
	// Método 1: Descargar y extraer .tar.gz
	libsail_tar_url := "https://github.com/HappySeaFox/sail/archive/refs/tags/v0.9.5.tar.gz"
	libsail_cmd_tar := exec.Command("curl", "-L", libsail_tar_url, "-o", "package.tar.gz")
	err := libsail_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libsail_zip_url := "https://github.com/HappySeaFox/sail/archive/refs/tags/v0.9.5.zip"
	libsail_cmd_zip := exec.Command("curl", "-L", libsail_zip_url, "-o", "package.zip")
	err = libsail_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libsail_bin_url := "https://github.com/HappySeaFox/sail/archive/refs/tags/v0.9.5.bin"
	libsail_cmd_bin := exec.Command("curl", "-L", libsail_bin_url, "-o", "binary.bin")
	err = libsail_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libsail_src_url := "https://github.com/HappySeaFox/sail/archive/refs/tags/v0.9.5.src.tar.gz"
	libsail_cmd_src := exec.Command("curl", "-L", libsail_src_url, "-o", "source.tar.gz")
	err = libsail_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libsail_cmd_direct := exec.Command("./binary")
	err = libsail_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: brotli")
	exec.Command("latte", "install", "brotli").Run()
	fmt.Println("Instalando dependencia: giflib")
	exec.Command("latte", "install", "giflib").Run()
	fmt.Println("Instalando dependencia: highway")
	exec.Command("latte", "install", "highway").Run()
	fmt.Println("Instalando dependencia: jasper")
	exec.Command("latte", "install", "jasper").Run()
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: jpeg-xl")
	exec.Command("latte", "install", "jpeg-xl").Run()
	fmt.Println("Instalando dependencia: libavif")
	exec.Command("latte", "install", "libavif").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: libtiff")
	exec.Command("latte", "install", "libtiff").Run()
	fmt.Println("Instalando dependencia: resvg")
	exec.Command("latte", "install", "resvg").Run()
	fmt.Println("Instalando dependencia: webp")
	exec.Command("latte", "install", "webp").Run()
}
