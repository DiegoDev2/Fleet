package main

// Pcalc - Calculator for those working with multiple bases, sizes, and close to the bits
// Homepage: https://github.com/alt-romes/programmer-calculator

import (
	"fmt"
	
	"os/exec"
)

func installPcalc() {
	// Método 1: Descargar y extraer .tar.gz
	pcalc_tar_url := "https://github.com/alt-romes/programmer-calculator/archive/refs/tags/v3.0.tar.gz"
	pcalc_cmd_tar := exec.Command("curl", "-L", pcalc_tar_url, "-o", "package.tar.gz")
	err := pcalc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pcalc_zip_url := "https://github.com/alt-romes/programmer-calculator/archive/refs/tags/v3.0.zip"
	pcalc_cmd_zip := exec.Command("curl", "-L", pcalc_zip_url, "-o", "package.zip")
	err = pcalc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pcalc_bin_url := "https://github.com/alt-romes/programmer-calculator/archive/refs/tags/v3.0.bin"
	pcalc_cmd_bin := exec.Command("curl", "-L", pcalc_bin_url, "-o", "binary.bin")
	err = pcalc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pcalc_src_url := "https://github.com/alt-romes/programmer-calculator/archive/refs/tags/v3.0.src.tar.gz"
	pcalc_cmd_src := exec.Command("curl", "-L", pcalc_src_url, "-o", "source.tar.gz")
	err = pcalc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pcalc_cmd_direct := exec.Command("./binary")
	err = pcalc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
