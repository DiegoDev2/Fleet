package main

// Meek - Blocking-resistant pluggable transport for Tor
// Homepage: https://www.torproject.org

import (
	"fmt"
	
	"os/exec"
)

func installMeek() {
	// Método 1: Descargar y extraer .tar.gz
	meek_tar_url := "https://gitlab.torproject.org/tpo/anti-censorship/pluggable-transports/meek/-/archive/v0.38.0/meek-v0.38.0.tar.gz"
	meek_cmd_tar := exec.Command("curl", "-L", meek_tar_url, "-o", "package.tar.gz")
	err := meek_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	meek_zip_url := "https://gitlab.torproject.org/tpo/anti-censorship/pluggable-transports/meek/-/archive/v0.38.0/meek-v0.38.0.zip"
	meek_cmd_zip := exec.Command("curl", "-L", meek_zip_url, "-o", "package.zip")
	err = meek_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	meek_bin_url := "https://gitlab.torproject.org/tpo/anti-censorship/pluggable-transports/meek/-/archive/v0.38.0/meek-v0.38.0.bin"
	meek_cmd_bin := exec.Command("curl", "-L", meek_bin_url, "-o", "binary.bin")
	err = meek_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	meek_src_url := "https://gitlab.torproject.org/tpo/anti-censorship/pluggable-transports/meek/-/archive/v0.38.0/meek-v0.38.0.src.tar.gz"
	meek_cmd_src := exec.Command("curl", "-L", meek_src_url, "-o", "source.tar.gz")
	err = meek_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	meek_cmd_direct := exec.Command("./binary")
	err = meek_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
