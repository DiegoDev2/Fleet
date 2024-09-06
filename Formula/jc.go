package main

// Jc - Serializes the output of command-line tools to structured JSON output
// Homepage: https://github.com/kellyjonbrazil/jc

import (
	"fmt"
	
	"os/exec"
)

func installJc() {
	// Método 1: Descargar y extraer .tar.gz
	jc_tar_url := "https://files.pythonhosted.org/packages/a5/82/bfb1ec7d9667bc2f922254bc62e12fd460a5de3b711518f5089df0a17180/jc-1.25.3.tar.gz"
	jc_cmd_tar := exec.Command("curl", "-L", jc_tar_url, "-o", "package.tar.gz")
	err := jc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jc_zip_url := "https://files.pythonhosted.org/packages/a5/82/bfb1ec7d9667bc2f922254bc62e12fd460a5de3b711518f5089df0a17180/jc-1.25.3.zip"
	jc_cmd_zip := exec.Command("curl", "-L", jc_zip_url, "-o", "package.zip")
	err = jc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jc_bin_url := "https://files.pythonhosted.org/packages/a5/82/bfb1ec7d9667bc2f922254bc62e12fd460a5de3b711518f5089df0a17180/jc-1.25.3.bin"
	jc_cmd_bin := exec.Command("curl", "-L", jc_bin_url, "-o", "binary.bin")
	err = jc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jc_src_url := "https://files.pythonhosted.org/packages/a5/82/bfb1ec7d9667bc2f922254bc62e12fd460a5de3b711518f5089df0a17180/jc-1.25.3.src.tar.gz"
	jc_cmd_src := exec.Command("curl", "-L", jc_src_url, "-o", "source.tar.gz")
	err = jc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jc_cmd_direct := exec.Command("./binary")
	err = jc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
