package main

// CyclonedxPython - Creates CycloneDX Software Bill of Materials (SBOM) from Python projects
// Homepage: https://cyclonedx.org/

import (
	"fmt"
	
	"os/exec"
)

func installCyclonedxPython() {
	// Método 1: Descargar y extraer .tar.gz
	cyclonedxpython_tar_url := "https://files.pythonhosted.org/packages/fd/19/1ea340c603724c36726ce562144eaf7517498564185bd978f59c50e57fbc/cyclonedx_bom-4.5.0.tar.gz"
	cyclonedxpython_cmd_tar := exec.Command("curl", "-L", cyclonedxpython_tar_url, "-o", "package.tar.gz")
	err := cyclonedxpython_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cyclonedxpython_zip_url := "https://files.pythonhosted.org/packages/fd/19/1ea340c603724c36726ce562144eaf7517498564185bd978f59c50e57fbc/cyclonedx_bom-4.5.0.zip"
	cyclonedxpython_cmd_zip := exec.Command("curl", "-L", cyclonedxpython_zip_url, "-o", "package.zip")
	err = cyclonedxpython_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cyclonedxpython_bin_url := "https://files.pythonhosted.org/packages/fd/19/1ea340c603724c36726ce562144eaf7517498564185bd978f59c50e57fbc/cyclonedx_bom-4.5.0.bin"
	cyclonedxpython_cmd_bin := exec.Command("curl", "-L", cyclonedxpython_bin_url, "-o", "binary.bin")
	err = cyclonedxpython_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cyclonedxpython_src_url := "https://files.pythonhosted.org/packages/fd/19/1ea340c603724c36726ce562144eaf7517498564185bd978f59c50e57fbc/cyclonedx_bom-4.5.0.src.tar.gz"
	cyclonedxpython_cmd_src := exec.Command("curl", "-L", cyclonedxpython_src_url, "-o", "source.tar.gz")
	err = cyclonedxpython_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cyclonedxpython_cmd_direct := exec.Command("./binary")
	err = cyclonedxpython_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
