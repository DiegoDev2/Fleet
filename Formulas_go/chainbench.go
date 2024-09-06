package main

// ChainBench - Software supply chain auditing tool based on CIS benchmark
// Homepage: https://github.com/aquasecurity/chain-bench

import (
	"fmt"
	
	"os/exec"
)

func installChainBench() {
	// Método 1: Descargar y extraer .tar.gz
	chainbench_tar_url := "https://github.com/aquasecurity/chain-bench/archive/refs/tags/v0.1.10.tar.gz"
	chainbench_cmd_tar := exec.Command("curl", "-L", chainbench_tar_url, "-o", "package.tar.gz")
	err := chainbench_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	chainbench_zip_url := "https://github.com/aquasecurity/chain-bench/archive/refs/tags/v0.1.10.zip"
	chainbench_cmd_zip := exec.Command("curl", "-L", chainbench_zip_url, "-o", "package.zip")
	err = chainbench_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	chainbench_bin_url := "https://github.com/aquasecurity/chain-bench/archive/refs/tags/v0.1.10.bin"
	chainbench_cmd_bin := exec.Command("curl", "-L", chainbench_bin_url, "-o", "binary.bin")
	err = chainbench_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	chainbench_src_url := "https://github.com/aquasecurity/chain-bench/archive/refs/tags/v0.1.10.src.tar.gz"
	chainbench_cmd_src := exec.Command("curl", "-L", chainbench_src_url, "-o", "source.tar.gz")
	err = chainbench_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	chainbench_cmd_direct := exec.Command("./binary")
	err = chainbench_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
