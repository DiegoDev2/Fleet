package main

// Drafter - Native C/C++ API Blueprint Parser
// Homepage: https://apiblueprint.org/

import (
	"fmt"
	
	"os/exec"
)

func installDrafter() {
	// Método 1: Descargar y extraer .tar.gz
	drafter_tar_url := "https://github.com/apiaryio/drafter/releases/download/v5.1.0/drafter-v5.1.0.tar.gz"
	drafter_cmd_tar := exec.Command("curl", "-L", drafter_tar_url, "-o", "package.tar.gz")
	err := drafter_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	drafter_zip_url := "https://github.com/apiaryio/drafter/releases/download/v5.1.0/drafter-v5.1.0.zip"
	drafter_cmd_zip := exec.Command("curl", "-L", drafter_zip_url, "-o", "package.zip")
	err = drafter_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	drafter_bin_url := "https://github.com/apiaryio/drafter/releases/download/v5.1.0/drafter-v5.1.0.bin"
	drafter_cmd_bin := exec.Command("curl", "-L", drafter_bin_url, "-o", "binary.bin")
	err = drafter_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	drafter_src_url := "https://github.com/apiaryio/drafter/releases/download/v5.1.0/drafter-v5.1.0.src.tar.gz"
	drafter_cmd_src := exec.Command("curl", "-L", drafter_src_url, "-o", "source.tar.gz")
	err = drafter_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	drafter_cmd_direct := exec.Command("./binary")
	err = drafter_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
