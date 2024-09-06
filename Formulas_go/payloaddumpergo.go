package main

// PayloadDumperGo - Android OTA payload dumper written in Go
// Homepage: https://github.com/ssut/payload-dumper-go

import (
	"fmt"
	
	"os/exec"
)

func installPayloadDumperGo() {
	// Método 1: Descargar y extraer .tar.gz
	payloaddumpergo_tar_url := "https://github.com/ssut/payload-dumper-go/archive/refs/tags/1.2.2.tar.gz"
	payloaddumpergo_cmd_tar := exec.Command("curl", "-L", payloaddumpergo_tar_url, "-o", "package.tar.gz")
	err := payloaddumpergo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	payloaddumpergo_zip_url := "https://github.com/ssut/payload-dumper-go/archive/refs/tags/1.2.2.zip"
	payloaddumpergo_cmd_zip := exec.Command("curl", "-L", payloaddumpergo_zip_url, "-o", "package.zip")
	err = payloaddumpergo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	payloaddumpergo_bin_url := "https://github.com/ssut/payload-dumper-go/archive/refs/tags/1.2.2.bin"
	payloaddumpergo_cmd_bin := exec.Command("curl", "-L", payloaddumpergo_bin_url, "-o", "binary.bin")
	err = payloaddumpergo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	payloaddumpergo_src_url := "https://github.com/ssut/payload-dumper-go/archive/refs/tags/1.2.2.src.tar.gz"
	payloaddumpergo_cmd_src := exec.Command("curl", "-L", payloaddumpergo_src_url, "-o", "source.tar.gz")
	err = payloaddumpergo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	payloaddumpergo_cmd_direct := exec.Command("./binary")
	err = payloaddumpergo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
	fmt.Println("Instalando dependencia: xz")
exec.Command("latte", "install", "xz")
}
