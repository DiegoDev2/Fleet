package main

// Minisign - Sign files & verify signatures. Works with signify in OpenBSD
// Homepage: https://jedisct1.github.io/minisign/

import (
	"fmt"
	
	"os/exec"
)

func installMinisign() {
	// Método 1: Descargar y extraer .tar.gz
	minisign_tar_url := "https://github.com/jedisct1/minisign/archive/refs/tags/0.11.tar.gz"
	minisign_cmd_tar := exec.Command("curl", "-L", minisign_tar_url, "-o", "package.tar.gz")
	err := minisign_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	minisign_zip_url := "https://github.com/jedisct1/minisign/archive/refs/tags/0.11.zip"
	minisign_cmd_zip := exec.Command("curl", "-L", minisign_zip_url, "-o", "package.zip")
	err = minisign_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	minisign_bin_url := "https://github.com/jedisct1/minisign/archive/refs/tags/0.11.bin"
	minisign_cmd_bin := exec.Command("curl", "-L", minisign_bin_url, "-o", "binary.bin")
	err = minisign_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	minisign_src_url := "https://github.com/jedisct1/minisign/archive/refs/tags/0.11.src.tar.gz"
	minisign_cmd_src := exec.Command("curl", "-L", minisign_src_url, "-o", "source.tar.gz")
	err = minisign_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	minisign_cmd_direct := exec.Command("./binary")
	err = minisign_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libsodium")
exec.Command("latte", "install", "libsodium")
}
