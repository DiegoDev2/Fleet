package main

// Numactl - NUMA support for Linux
// Homepage: https://github.com/numactl/numactl

import (
	"fmt"
	
	"os/exec"
)

func installNumactl() {
	// Método 1: Descargar y extraer .tar.gz
	numactl_tar_url := "https://github.com/numactl/numactl/releases/download/v2.0.18/numactl-2.0.18.tar.gz"
	numactl_cmd_tar := exec.Command("curl", "-L", numactl_tar_url, "-o", "package.tar.gz")
	err := numactl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	numactl_zip_url := "https://github.com/numactl/numactl/releases/download/v2.0.18/numactl-2.0.18.zip"
	numactl_cmd_zip := exec.Command("curl", "-L", numactl_zip_url, "-o", "package.zip")
	err = numactl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	numactl_bin_url := "https://github.com/numactl/numactl/releases/download/v2.0.18/numactl-2.0.18.bin"
	numactl_cmd_bin := exec.Command("curl", "-L", numactl_bin_url, "-o", "binary.bin")
	err = numactl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	numactl_src_url := "https://github.com/numactl/numactl/releases/download/v2.0.18/numactl-2.0.18.src.tar.gz"
	numactl_cmd_src := exec.Command("curl", "-L", numactl_src_url, "-o", "source.tar.gz")
	err = numactl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	numactl_cmd_direct := exec.Command("./binary")
	err = numactl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
