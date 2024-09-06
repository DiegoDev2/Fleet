package main

// Safeint - Class library for C++ that manages integer overflows
// Homepage: https://github.com/dcleblanc/SafeInt

import (
	"fmt"
	
	"os/exec"
)

func installSafeint() {
	// Método 1: Descargar y extraer .tar.gz
	safeint_tar_url := "https://github.com/dcleblanc/SafeInt/archive/refs/tags/3.0.28a.tar.gz"
	safeint_cmd_tar := exec.Command("curl", "-L", safeint_tar_url, "-o", "package.tar.gz")
	err := safeint_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	safeint_zip_url := "https://github.com/dcleblanc/SafeInt/archive/refs/tags/3.0.28a.zip"
	safeint_cmd_zip := exec.Command("curl", "-L", safeint_zip_url, "-o", "package.zip")
	err = safeint_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	safeint_bin_url := "https://github.com/dcleblanc/SafeInt/archive/refs/tags/3.0.28a.bin"
	safeint_cmd_bin := exec.Command("curl", "-L", safeint_bin_url, "-o", "binary.bin")
	err = safeint_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	safeint_src_url := "https://github.com/dcleblanc/SafeInt/archive/refs/tags/3.0.28a.src.tar.gz"
	safeint_cmd_src := exec.Command("curl", "-L", safeint_src_url, "-o", "source.tar.gz")
	err = safeint_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	safeint_cmd_direct := exec.Command("./binary")
	err = safeint_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
