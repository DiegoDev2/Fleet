package main

// Traefik - Modern reverse proxy
// Homepage: https://traefik.io/

import (
	"fmt"
	
	"os/exec"
)

func installTraefik() {
	// Método 1: Descargar y extraer .tar.gz
	traefik_tar_url := "https://github.com/traefik/traefik/releases/download/v3.1.2/traefik-v3.1.2.src.tar.gz"
	traefik_cmd_tar := exec.Command("curl", "-L", traefik_tar_url, "-o", "package.tar.gz")
	err := traefik_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	traefik_zip_url := "https://github.com/traefik/traefik/releases/download/v3.1.2/traefik-v3.1.2.src.zip"
	traefik_cmd_zip := exec.Command("curl", "-L", traefik_zip_url, "-o", "package.zip")
	err = traefik_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	traefik_bin_url := "https://github.com/traefik/traefik/releases/download/v3.1.2/traefik-v3.1.2.src.bin"
	traefik_cmd_bin := exec.Command("curl", "-L", traefik_bin_url, "-o", "binary.bin")
	err = traefik_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	traefik_src_url := "https://github.com/traefik/traefik/releases/download/v3.1.2/traefik-v3.1.2.src.src.tar.gz"
	traefik_cmd_src := exec.Command("curl", "-L", traefik_src_url, "-o", "source.tar.gz")
	err = traefik_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	traefik_cmd_direct := exec.Command("./binary")
	err = traefik_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
