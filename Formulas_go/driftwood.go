package main

// Driftwood - Private key usage verification
// Homepage: https://trufflesecurity.com/

import (
	"fmt"
	
	"os/exec"
)

func installDriftwood() {
	// Método 1: Descargar y extraer .tar.gz
	driftwood_tar_url := "https://github.com/trufflesecurity/driftwood/archive/refs/tags/v1.0.1.tar.gz"
	driftwood_cmd_tar := exec.Command("curl", "-L", driftwood_tar_url, "-o", "package.tar.gz")
	err := driftwood_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	driftwood_zip_url := "https://github.com/trufflesecurity/driftwood/archive/refs/tags/v1.0.1.zip"
	driftwood_cmd_zip := exec.Command("curl", "-L", driftwood_zip_url, "-o", "package.zip")
	err = driftwood_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	driftwood_bin_url := "https://github.com/trufflesecurity/driftwood/archive/refs/tags/v1.0.1.bin"
	driftwood_cmd_bin := exec.Command("curl", "-L", driftwood_bin_url, "-o", "binary.bin")
	err = driftwood_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	driftwood_src_url := "https://github.com/trufflesecurity/driftwood/archive/refs/tags/v1.0.1.src.tar.gz"
	driftwood_cmd_src := exec.Command("curl", "-L", driftwood_src_url, "-o", "source.tar.gz")
	err = driftwood_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	driftwood_cmd_direct := exec.Command("./binary")
	err = driftwood_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
