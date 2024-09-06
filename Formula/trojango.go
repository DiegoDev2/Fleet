package main

// TrojanGo - Trojan proxy in Go
// Homepage: https://p4gefau1t.github.io/trojan-go/

import (
	"fmt"
	
	"os/exec"
)

func installTrojanGo() {
	// Método 1: Descargar y extraer .tar.gz
	trojango_tar_url := "https://github.com/p4gefau1t/trojan-go.git"
	trojango_cmd_tar := exec.Command("curl", "-L", trojango_tar_url, "-o", "package.tar.gz")
	err := trojango_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	trojango_zip_url := "https://github.com/p4gefau1t/trojan-go.git"
	trojango_cmd_zip := exec.Command("curl", "-L", trojango_zip_url, "-o", "package.zip")
	err = trojango_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	trojango_bin_url := "https://github.com/p4gefau1t/trojan-go.git"
	trojango_cmd_bin := exec.Command("curl", "-L", trojango_bin_url, "-o", "binary.bin")
	err = trojango_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	trojango_src_url := "https://github.com/p4gefau1t/trojan-go.git"
	trojango_cmd_src := exec.Command("curl", "-L", trojango_src_url, "-o", "source.tar.gz")
	err = trojango_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	trojango_cmd_direct := exec.Command("./binary")
	err = trojango_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
