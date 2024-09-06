package main

// FuzzyFind - Fuzzy filename finder matching across directories as well as files
// Homepage: https://github.com/silentbicycle/ff

import (
	"fmt"
	
	"os/exec"
)

func installFuzzyFind() {
	// Método 1: Descargar y extraer .tar.gz
	fuzzyfind_tar_url := "https://github.com/silentbicycle/ff/archive/refs/tags/v0.6-flag-features.tar.gz"
	fuzzyfind_cmd_tar := exec.Command("curl", "-L", fuzzyfind_tar_url, "-o", "package.tar.gz")
	err := fuzzyfind_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fuzzyfind_zip_url := "https://github.com/silentbicycle/ff/archive/refs/tags/v0.6-flag-features.zip"
	fuzzyfind_cmd_zip := exec.Command("curl", "-L", fuzzyfind_zip_url, "-o", "package.zip")
	err = fuzzyfind_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fuzzyfind_bin_url := "https://github.com/silentbicycle/ff/archive/refs/tags/v0.6-flag-features.bin"
	fuzzyfind_cmd_bin := exec.Command("curl", "-L", fuzzyfind_bin_url, "-o", "binary.bin")
	err = fuzzyfind_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fuzzyfind_src_url := "https://github.com/silentbicycle/ff/archive/refs/tags/v0.6-flag-features.src.tar.gz"
	fuzzyfind_cmd_src := exec.Command("curl", "-L", fuzzyfind_src_url, "-o", "source.tar.gz")
	err = fuzzyfind_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fuzzyfind_cmd_direct := exec.Command("./binary")
	err = fuzzyfind_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
