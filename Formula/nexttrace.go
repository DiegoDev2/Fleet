package main

// Nexttrace - Open source visual route tracking CLI tool
// Homepage: https://nxtrace.github.io/NTrace-core/

import (
	"fmt"
	
	"os/exec"
)

func installNexttrace() {
	// Método 1: Descargar y extraer .tar.gz
	nexttrace_tar_url := "https://github.com/nxtrace/NTrace-core/archive/refs/tags/v1.3.2.tar.gz"
	nexttrace_cmd_tar := exec.Command("curl", "-L", nexttrace_tar_url, "-o", "package.tar.gz")
	err := nexttrace_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nexttrace_zip_url := "https://github.com/nxtrace/NTrace-core/archive/refs/tags/v1.3.2.zip"
	nexttrace_cmd_zip := exec.Command("curl", "-L", nexttrace_zip_url, "-o", "package.zip")
	err = nexttrace_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nexttrace_bin_url := "https://github.com/nxtrace/NTrace-core/archive/refs/tags/v1.3.2.bin"
	nexttrace_cmd_bin := exec.Command("curl", "-L", nexttrace_bin_url, "-o", "binary.bin")
	err = nexttrace_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nexttrace_src_url := "https://github.com/nxtrace/NTrace-core/archive/refs/tags/v1.3.2.src.tar.gz"
	nexttrace_cmd_src := exec.Command("curl", "-L", nexttrace_src_url, "-o", "source.tar.gz")
	err = nexttrace_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nexttrace_cmd_direct := exec.Command("./binary")
	err = nexttrace_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go@1.22")
	exec.Command("latte", "install", "go@1.22").Run()
}
