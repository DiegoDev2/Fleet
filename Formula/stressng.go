package main

// StressNg - Stress test a computer system in various selectable ways
// Homepage: https://wiki.ubuntu.com/Kernel/Reference/stress-ng

import (
	"fmt"
	
	"os/exec"
)

func installStressNg() {
	// Método 1: Descargar y extraer .tar.gz
	stressng_tar_url := "https://github.com/ColinIanKing/stress-ng/archive/refs/tags/V0.18.03.tar.gz"
	stressng_cmd_tar := exec.Command("curl", "-L", stressng_tar_url, "-o", "package.tar.gz")
	err := stressng_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	stressng_zip_url := "https://github.com/ColinIanKing/stress-ng/archive/refs/tags/V0.18.03.zip"
	stressng_cmd_zip := exec.Command("curl", "-L", stressng_zip_url, "-o", "package.zip")
	err = stressng_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	stressng_bin_url := "https://github.com/ColinIanKing/stress-ng/archive/refs/tags/V0.18.03.bin"
	stressng_cmd_bin := exec.Command("curl", "-L", stressng_bin_url, "-o", "binary.bin")
	err = stressng_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	stressng_src_url := "https://github.com/ColinIanKing/stress-ng/archive/refs/tags/V0.18.03.src.tar.gz"
	stressng_cmd_src := exec.Command("curl", "-L", stressng_src_url, "-o", "source.tar.gz")
	err = stressng_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	stressng_cmd_direct := exec.Command("./binary")
	err = stressng_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
