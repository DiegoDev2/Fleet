package main

// Libkeccak - Keccak-family hashing library
// Homepage: https://codeberg.org/maandree/libkeccak

import (
	"fmt"
	
	"os/exec"
)

func installLibkeccak() {
	// Método 1: Descargar y extraer .tar.gz
	libkeccak_tar_url := "https://codeberg.org/maandree/libkeccak/archive/1.4.tar.gz"
	libkeccak_cmd_tar := exec.Command("curl", "-L", libkeccak_tar_url, "-o", "package.tar.gz")
	err := libkeccak_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libkeccak_zip_url := "https://codeberg.org/maandree/libkeccak/archive/1.4.zip"
	libkeccak_cmd_zip := exec.Command("curl", "-L", libkeccak_zip_url, "-o", "package.zip")
	err = libkeccak_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libkeccak_bin_url := "https://codeberg.org/maandree/libkeccak/archive/1.4.bin"
	libkeccak_cmd_bin := exec.Command("curl", "-L", libkeccak_bin_url, "-o", "binary.bin")
	err = libkeccak_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libkeccak_src_url := "https://codeberg.org/maandree/libkeccak/archive/1.4.src.tar.gz"
	libkeccak_cmd_src := exec.Command("curl", "-L", libkeccak_src_url, "-o", "source.tar.gz")
	err = libkeccak_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libkeccak_cmd_direct := exec.Command("./binary")
	err = libkeccak_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
