package main

// Trailscraper - Tool to get valuable information out of AWS CloudTrail
// Homepage: https://github.com/flosell/trailscraper

import (
	"fmt"
	
	"os/exec"
)

func installTrailscraper() {
	// Método 1: Descargar y extraer .tar.gz
	trailscraper_tar_url := "https://files.pythonhosted.org/packages/bc/9b/f425ff02c84a16e434526d3ffe7abfc50589f46a5efe9b02cfd09bec698e/trailscraper-0.8.1.tar.gz"
	trailscraper_cmd_tar := exec.Command("curl", "-L", trailscraper_tar_url, "-o", "package.tar.gz")
	err := trailscraper_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	trailscraper_zip_url := "https://files.pythonhosted.org/packages/bc/9b/f425ff02c84a16e434526d3ffe7abfc50589f46a5efe9b02cfd09bec698e/trailscraper-0.8.1.zip"
	trailscraper_cmd_zip := exec.Command("curl", "-L", trailscraper_zip_url, "-o", "package.zip")
	err = trailscraper_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	trailscraper_bin_url := "https://files.pythonhosted.org/packages/bc/9b/f425ff02c84a16e434526d3ffe7abfc50589f46a5efe9b02cfd09bec698e/trailscraper-0.8.1.bin"
	trailscraper_cmd_bin := exec.Command("curl", "-L", trailscraper_bin_url, "-o", "binary.bin")
	err = trailscraper_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	trailscraper_src_url := "https://files.pythonhosted.org/packages/bc/9b/f425ff02c84a16e434526d3ffe7abfc50589f46a5efe9b02cfd09bec698e/trailscraper-0.8.1.src.tar.gz"
	trailscraper_cmd_src := exec.Command("curl", "-L", trailscraper_src_url, "-o", "source.tar.gz")
	err = trailscraper_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	trailscraper_cmd_direct := exec.Command("./binary")
	err = trailscraper_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
