package main

// Bupstash - Easy and efficient encrypted backups
// Homepage: https://bupstash.io

import (
	"fmt"
	
	"os/exec"
)

func installBupstash() {
	// Método 1: Descargar y extraer .tar.gz
	bupstash_tar_url := "https://github.com/andrewchambers/bupstash/releases/download/v0.12.0/bupstash-v0.12.0-src+deps.tar.gz"
	bupstash_cmd_tar := exec.Command("curl", "-L", bupstash_tar_url, "-o", "package.tar.gz")
	err := bupstash_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bupstash_zip_url := "https://github.com/andrewchambers/bupstash/releases/download/v0.12.0/bupstash-v0.12.0-src+deps.zip"
	bupstash_cmd_zip := exec.Command("curl", "-L", bupstash_zip_url, "-o", "package.zip")
	err = bupstash_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bupstash_bin_url := "https://github.com/andrewchambers/bupstash/releases/download/v0.12.0/bupstash-v0.12.0-src+deps.bin"
	bupstash_cmd_bin := exec.Command("curl", "-L", bupstash_bin_url, "-o", "binary.bin")
	err = bupstash_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bupstash_src_url := "https://github.com/andrewchambers/bupstash/releases/download/v0.12.0/bupstash-v0.12.0-src+deps.src.tar.gz"
	bupstash_cmd_src := exec.Command("curl", "-L", bupstash_src_url, "-o", "source.tar.gz")
	err = bupstash_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bupstash_cmd_direct := exec.Command("./binary")
	err = bupstash_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: libsodium")
exec.Command("latte", "install", "libsodium")
}
