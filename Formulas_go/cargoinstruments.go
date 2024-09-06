package main

// CargoInstruments - Easily generate Instruments traces for your rust crate
// Homepage: https://github.com/cmyr/cargo-instruments

import (
	"fmt"
	
	"os/exec"
)

func installCargoInstruments() {
	// Método 1: Descargar y extraer .tar.gz
	cargoinstruments_tar_url := "https://github.com/cmyr/cargo-instruments/archive/refs/tags/v0.4.10.tar.gz"
	cargoinstruments_cmd_tar := exec.Command("curl", "-L", cargoinstruments_tar_url, "-o", "package.tar.gz")
	err := cargoinstruments_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cargoinstruments_zip_url := "https://github.com/cmyr/cargo-instruments/archive/refs/tags/v0.4.10.zip"
	cargoinstruments_cmd_zip := exec.Command("curl", "-L", cargoinstruments_zip_url, "-o", "package.zip")
	err = cargoinstruments_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cargoinstruments_bin_url := "https://github.com/cmyr/cargo-instruments/archive/refs/tags/v0.4.10.bin"
	cargoinstruments_cmd_bin := exec.Command("curl", "-L", cargoinstruments_bin_url, "-o", "binary.bin")
	err = cargoinstruments_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cargoinstruments_src_url := "https://github.com/cmyr/cargo-instruments/archive/refs/tags/v0.4.10.src.tar.gz"
	cargoinstruments_cmd_src := exec.Command("curl", "-L", cargoinstruments_src_url, "-o", "source.tar.gz")
	err = cargoinstruments_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cargoinstruments_cmd_direct := exec.Command("./binary")
	err = cargoinstruments_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
