package main

// Skopeo - Work with remote images registries
// Homepage: https://github.com/containers/skopeo

import (
	"fmt"
	
	"os/exec"
)

func installSkopeo() {
	// Método 1: Descargar y extraer .tar.gz
	skopeo_tar_url := "https://github.com/containers/skopeo/archive/refs/tags/v1.16.1.tar.gz"
	skopeo_cmd_tar := exec.Command("curl", "-L", skopeo_tar_url, "-o", "package.tar.gz")
	err := skopeo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	skopeo_zip_url := "https://github.com/containers/skopeo/archive/refs/tags/v1.16.1.zip"
	skopeo_cmd_zip := exec.Command("curl", "-L", skopeo_zip_url, "-o", "package.zip")
	err = skopeo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	skopeo_bin_url := "https://github.com/containers/skopeo/archive/refs/tags/v1.16.1.bin"
	skopeo_cmd_bin := exec.Command("curl", "-L", skopeo_bin_url, "-o", "binary.bin")
	err = skopeo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	skopeo_src_url := "https://github.com/containers/skopeo/archive/refs/tags/v1.16.1.src.tar.gz"
	skopeo_cmd_src := exec.Command("curl", "-L", skopeo_src_url, "-o", "source.tar.gz")
	err = skopeo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	skopeo_cmd_direct := exec.Command("./binary")
	err = skopeo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
	fmt.Println("Instalando dependencia: go-md2man")
exec.Command("latte", "install", "go-md2man")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: gpgme")
exec.Command("latte", "install", "gpgme")
	fmt.Println("Instalando dependencia: device-mapper")
exec.Command("latte", "install", "device-mapper")
}
