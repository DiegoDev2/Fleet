package main

// VampPluginSdk - Audio processing plugin system sdk
// Homepage: https://www.vamp-plugins.org/

import (
	"fmt"
	
	"os/exec"
)

func installVampPluginSdk() {
	// Método 1: Descargar y extraer .tar.gz
	vamppluginsdk_tar_url := "https://deb.debian.org/debian/pool/main/v/vamp-plugin-sdk/vamp-plugin-sdk_2.10.0.orig.tar.gz"
	vamppluginsdk_cmd_tar := exec.Command("curl", "-L", vamppluginsdk_tar_url, "-o", "package.tar.gz")
	err := vamppluginsdk_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vamppluginsdk_zip_url := "https://deb.debian.org/debian/pool/main/v/vamp-plugin-sdk/vamp-plugin-sdk_2.10.0.orig.zip"
	vamppluginsdk_cmd_zip := exec.Command("curl", "-L", vamppluginsdk_zip_url, "-o", "package.zip")
	err = vamppluginsdk_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vamppluginsdk_bin_url := "https://deb.debian.org/debian/pool/main/v/vamp-plugin-sdk/vamp-plugin-sdk_2.10.0.orig.bin"
	vamppluginsdk_cmd_bin := exec.Command("curl", "-L", vamppluginsdk_bin_url, "-o", "binary.bin")
	err = vamppluginsdk_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vamppluginsdk_src_url := "https://deb.debian.org/debian/pool/main/v/vamp-plugin-sdk/vamp-plugin-sdk_2.10.0.orig.src.tar.gz"
	vamppluginsdk_cmd_src := exec.Command("curl", "-L", vamppluginsdk_src_url, "-o", "source.tar.gz")
	err = vamppluginsdk_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vamppluginsdk_cmd_direct := exec.Command("./binary")
	err = vamppluginsdk_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: flac")
exec.Command("latte", "install", "flac")
	fmt.Println("Instalando dependencia: libogg")
exec.Command("latte", "install", "libogg")
	fmt.Println("Instalando dependencia: libsndfile")
exec.Command("latte", "install", "libsndfile")
}
