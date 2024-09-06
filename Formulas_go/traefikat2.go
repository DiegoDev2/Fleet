package main

// TraefikAT2 - Modern reverse proxy
// Homepage: https://traefik.io/

import (
	"fmt"
	
	"os/exec"
)

func installTraefikAT2() {
	// Método 1: Descargar y extraer .tar.gz
	traefikat2_tar_url := "https://github.com/traefik/traefik/releases/download/v2.11.8/traefik-v2.11.8.src.tar.gz"
	traefikat2_cmd_tar := exec.Command("curl", "-L", traefikat2_tar_url, "-o", "package.tar.gz")
	err := traefikat2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	traefikat2_zip_url := "https://github.com/traefik/traefik/releases/download/v2.11.8/traefik-v2.11.8.src.zip"
	traefikat2_cmd_zip := exec.Command("curl", "-L", traefikat2_zip_url, "-o", "package.zip")
	err = traefikat2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	traefikat2_bin_url := "https://github.com/traefik/traefik/releases/download/v2.11.8/traefik-v2.11.8.src.bin"
	traefikat2_cmd_bin := exec.Command("curl", "-L", traefikat2_bin_url, "-o", "binary.bin")
	err = traefikat2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	traefikat2_src_url := "https://github.com/traefik/traefik/releases/download/v2.11.8/traefik-v2.11.8.src.src.tar.gz"
	traefikat2_cmd_src := exec.Command("curl", "-L", traefikat2_src_url, "-o", "source.tar.gz")
	err = traefikat2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	traefikat2_cmd_direct := exec.Command("./binary")
	err = traefikat2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
