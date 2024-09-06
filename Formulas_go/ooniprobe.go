package main

// Ooniprobe - Network interference detection tool
// Homepage: https://ooni.org/

import (
	"fmt"
	
	"os/exec"
)

func installOoniprobe() {
	// Método 1: Descargar y extraer .tar.gz
	ooniprobe_tar_url := "https://github.com/ooni/probe-cli/archive/refs/tags/v3.23.0.tar.gz"
	ooniprobe_cmd_tar := exec.Command("curl", "-L", ooniprobe_tar_url, "-o", "package.tar.gz")
	err := ooniprobe_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ooniprobe_zip_url := "https://github.com/ooni/probe-cli/archive/refs/tags/v3.23.0.zip"
	ooniprobe_cmd_zip := exec.Command("curl", "-L", ooniprobe_zip_url, "-o", "package.zip")
	err = ooniprobe_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ooniprobe_bin_url := "https://github.com/ooni/probe-cli/archive/refs/tags/v3.23.0.bin"
	ooniprobe_cmd_bin := exec.Command("curl", "-L", ooniprobe_bin_url, "-o", "binary.bin")
	err = ooniprobe_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ooniprobe_src_url := "https://github.com/ooni/probe-cli/archive/refs/tags/v3.23.0.src.tar.gz"
	ooniprobe_cmd_src := exec.Command("curl", "-L", ooniprobe_src_url, "-o", "source.tar.gz")
	err = ooniprobe_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ooniprobe_cmd_direct := exec.Command("./binary")
	err = ooniprobe_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
	fmt.Println("Instalando dependencia: tor")
exec.Command("latte", "install", "tor")
}
