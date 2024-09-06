package main

// Kvazaar - Ultravideo HEVC encoder
// Homepage: https://github.com/ultravideo/kvazaar

import (
	"fmt"
	
	"os/exec"
)

func installKvazaar() {
	// Método 1: Descargar y extraer .tar.gz
	kvazaar_tar_url := "https://github.com/ultravideo/kvazaar/archive/refs/tags/v2.3.1.tar.gz"
	kvazaar_cmd_tar := exec.Command("curl", "-L", kvazaar_tar_url, "-o", "package.tar.gz")
	err := kvazaar_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kvazaar_zip_url := "https://github.com/ultravideo/kvazaar/archive/refs/tags/v2.3.1.zip"
	kvazaar_cmd_zip := exec.Command("curl", "-L", kvazaar_zip_url, "-o", "package.zip")
	err = kvazaar_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kvazaar_bin_url := "https://github.com/ultravideo/kvazaar/archive/refs/tags/v2.3.1.bin"
	kvazaar_cmd_bin := exec.Command("curl", "-L", kvazaar_bin_url, "-o", "binary.bin")
	err = kvazaar_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kvazaar_src_url := "https://github.com/ultravideo/kvazaar/archive/refs/tags/v2.3.1.src.tar.gz"
	kvazaar_cmd_src := exec.Command("curl", "-L", kvazaar_src_url, "-o", "source.tar.gz")
	err = kvazaar_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kvazaar_cmd_direct := exec.Command("./binary")
	err = kvazaar_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: yasm")
exec.Command("latte", "install", "yasm")
}
