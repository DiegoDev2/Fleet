package main

// Opendht - C++17 Distributed Hash Table implementation
// Homepage: https://github.com/savoirfairelinux/opendht

import (
	"fmt"
	
	"os/exec"
)

func installOpendht() {
	// Método 1: Descargar y extraer .tar.gz
	opendht_tar_url := "https://github.com/savoirfairelinux/opendht/archive/refs/tags/v3.2.0.tar.gz"
	opendht_cmd_tar := exec.Command("curl", "-L", opendht_tar_url, "-o", "package.tar.gz")
	err := opendht_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	opendht_zip_url := "https://github.com/savoirfairelinux/opendht/archive/refs/tags/v3.2.0.zip"
	opendht_cmd_zip := exec.Command("curl", "-L", opendht_zip_url, "-o", "package.zip")
	err = opendht_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	opendht_bin_url := "https://github.com/savoirfairelinux/opendht/archive/refs/tags/v3.2.0.bin"
	opendht_cmd_bin := exec.Command("curl", "-L", opendht_bin_url, "-o", "binary.bin")
	err = opendht_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	opendht_src_url := "https://github.com/savoirfairelinux/opendht/archive/refs/tags/v3.2.0.src.tar.gz"
	opendht_cmd_src := exec.Command("curl", "-L", opendht_src_url, "-o", "source.tar.gz")
	err = opendht_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	opendht_cmd_direct := exec.Command("./binary")
	err = opendht_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: argon2")
	exec.Command("latte", "install", "argon2").Run()
	fmt.Println("Instalando dependencia: asio")
	exec.Command("latte", "install", "asio").Run()
	fmt.Println("Instalando dependencia: fmt")
	exec.Command("latte", "install", "fmt").Run()
	fmt.Println("Instalando dependencia: gnutls")
	exec.Command("latte", "install", "gnutls").Run()
	fmt.Println("Instalando dependencia: msgpack-cxx")
	exec.Command("latte", "install", "msgpack-cxx").Run()
	fmt.Println("Instalando dependencia: nettle")
	exec.Command("latte", "install", "nettle").Run()
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
}
