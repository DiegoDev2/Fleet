package main

// Keydb - Multithreaded fork of Redis
// Homepage: https://keydb.dev

import (
	"fmt"
	
	"os/exec"
)

func installKeydb() {
	// Método 1: Descargar y extraer .tar.gz
	keydb_tar_url := "https://github.com/Snapchat/KeyDB/archive/refs/tags/v6.3.4.tar.gz"
	keydb_cmd_tar := exec.Command("curl", "-L", keydb_tar_url, "-o", "package.tar.gz")
	err := keydb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	keydb_zip_url := "https://github.com/Snapchat/KeyDB/archive/refs/tags/v6.3.4.zip"
	keydb_cmd_zip := exec.Command("curl", "-L", keydb_zip_url, "-o", "package.zip")
	err = keydb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	keydb_bin_url := "https://github.com/Snapchat/KeyDB/archive/refs/tags/v6.3.4.bin"
	keydb_cmd_bin := exec.Command("curl", "-L", keydb_bin_url, "-o", "binary.bin")
	err = keydb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	keydb_src_url := "https://github.com/Snapchat/KeyDB/archive/refs/tags/v6.3.4.src.tar.gz"
	keydb_cmd_src := exec.Command("curl", "-L", keydb_src_url, "-o", "source.tar.gz")
	err = keydb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	keydb_cmd_direct := exec.Command("./binary")
	err = keydb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: snappy")
exec.Command("latte", "install", "snappy")
	fmt.Println("Instalando dependencia: zstd")
exec.Command("latte", "install", "zstd")
	fmt.Println("Instalando dependencia: util-linux")
exec.Command("latte", "install", "util-linux")
}
