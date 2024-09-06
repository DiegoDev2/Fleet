package main

// Epr - Command-line EPUB reader
// Homepage: https://github.com/wustho/epr

import (
	"fmt"
	
	"os/exec"
)

func installEpr() {
	// Método 1: Descargar y extraer .tar.gz
	epr_tar_url := "https://files.pythonhosted.org/packages/39/20/d647083aa86ec9da89b4f04b62dd6942aabb77528fd2efe018ff1cd145d2/epr-reader-2.4.15.tar.gz"
	epr_cmd_tar := exec.Command("curl", "-L", epr_tar_url, "-o", "package.tar.gz")
	err := epr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	epr_zip_url := "https://files.pythonhosted.org/packages/39/20/d647083aa86ec9da89b4f04b62dd6942aabb77528fd2efe018ff1cd145d2/epr-reader-2.4.15.zip"
	epr_cmd_zip := exec.Command("curl", "-L", epr_zip_url, "-o", "package.zip")
	err = epr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	epr_bin_url := "https://files.pythonhosted.org/packages/39/20/d647083aa86ec9da89b4f04b62dd6942aabb77528fd2efe018ff1cd145d2/epr-reader-2.4.15.bin"
	epr_cmd_bin := exec.Command("curl", "-L", epr_bin_url, "-o", "binary.bin")
	err = epr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	epr_src_url := "https://files.pythonhosted.org/packages/39/20/d647083aa86ec9da89b4f04b62dd6942aabb77528fd2efe018ff1cd145d2/epr-reader-2.4.15.src.tar.gz"
	epr_cmd_src := exec.Command("curl", "-L", epr_src_url, "-o", "source.tar.gz")
	err = epr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	epr_cmd_direct := exec.Command("./binary")
	err = epr_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
