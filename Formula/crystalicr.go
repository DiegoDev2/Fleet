package main

// CrystalIcr - Interactive console for Crystal programming language
// Homepage: https://github.com/crystal-community/icr

import (
	"fmt"
	
	"os/exec"
)

func installCrystalIcr() {
	// Método 1: Descargar y extraer .tar.gz
	crystalicr_tar_url := "https://github.com/crystal-community/icr/archive/refs/tags/v0.9.0.tar.gz"
	crystalicr_cmd_tar := exec.Command("curl", "-L", crystalicr_tar_url, "-o", "package.tar.gz")
	err := crystalicr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	crystalicr_zip_url := "https://github.com/crystal-community/icr/archive/refs/tags/v0.9.0.zip"
	crystalicr_cmd_zip := exec.Command("curl", "-L", crystalicr_zip_url, "-o", "package.zip")
	err = crystalicr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	crystalicr_bin_url := "https://github.com/crystal-community/icr/archive/refs/tags/v0.9.0.bin"
	crystalicr_cmd_bin := exec.Command("curl", "-L", crystalicr_bin_url, "-o", "binary.bin")
	err = crystalicr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	crystalicr_src_url := "https://github.com/crystal-community/icr/archive/refs/tags/v0.9.0.src.tar.gz"
	crystalicr_cmd_src := exec.Command("curl", "-L", crystalicr_src_url, "-o", "source.tar.gz")
	err = crystalicr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	crystalicr_cmd_direct := exec.Command("./binary")
	err = crystalicr_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bdw-gc")
	exec.Command("latte", "install", "bdw-gc").Run()
	fmt.Println("Instalando dependencia: crystal")
	exec.Command("latte", "install", "crystal").Run()
	fmt.Println("Instalando dependencia: libevent")
	exec.Command("latte", "install", "libevent").Run()
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: pcre2")
	exec.Command("latte", "install", "pcre2").Run()
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
}
