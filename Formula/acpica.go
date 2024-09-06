package main

// Acpica - OS-independent implementation of the ACPI specification
// Homepage: https://www.intel.com/content/www/us/en/developer/topic-technology/open/acpica/overview.html

import (
	"fmt"
	
	"os/exec"
)

func installAcpica() {
	// Método 1: Descargar y extraer .tar.gz
	acpica_tar_url := "https://downloadmirror.intel.com/819451/acpica-unix-20240321.tar.gz"
	acpica_cmd_tar := exec.Command("curl", "-L", acpica_tar_url, "-o", "package.tar.gz")
	err := acpica_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	acpica_zip_url := "https://downloadmirror.intel.com/819451/acpica-unix-20240321.zip"
	acpica_cmd_zip := exec.Command("curl", "-L", acpica_zip_url, "-o", "package.zip")
	err = acpica_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	acpica_bin_url := "https://downloadmirror.intel.com/819451/acpica-unix-20240321.bin"
	acpica_cmd_bin := exec.Command("curl", "-L", acpica_bin_url, "-o", "binary.bin")
	err = acpica_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	acpica_src_url := "https://downloadmirror.intel.com/819451/acpica-unix-20240321.src.tar.gz"
	acpica_cmd_src := exec.Command("curl", "-L", acpica_src_url, "-o", "source.tar.gz")
	err = acpica_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	acpica_cmd_direct := exec.Command("./binary")
	err = acpica_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
