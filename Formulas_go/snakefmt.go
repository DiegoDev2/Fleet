package main

// Snakefmt - Snakemake code formatter
// Homepage: https://github.com/snakemake/snakefmt/

import (
	"fmt"
	
	"os/exec"
)

func installSnakefmt() {
	// Método 1: Descargar y extraer .tar.gz
	snakefmt_tar_url := "https://files.pythonhosted.org/packages/b4/61/0228586a10b76064431e1d0c965f030b4c7dfbea6d1cfcb4d3f0cb0e6726/snakefmt-0.10.2.tar.gz"
	snakefmt_cmd_tar := exec.Command("curl", "-L", snakefmt_tar_url, "-o", "package.tar.gz")
	err := snakefmt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	snakefmt_zip_url := "https://files.pythonhosted.org/packages/b4/61/0228586a10b76064431e1d0c965f030b4c7dfbea6d1cfcb4d3f0cb0e6726/snakefmt-0.10.2.zip"
	snakefmt_cmd_zip := exec.Command("curl", "-L", snakefmt_zip_url, "-o", "package.zip")
	err = snakefmt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	snakefmt_bin_url := "https://files.pythonhosted.org/packages/b4/61/0228586a10b76064431e1d0c965f030b4c7dfbea6d1cfcb4d3f0cb0e6726/snakefmt-0.10.2.bin"
	snakefmt_cmd_bin := exec.Command("curl", "-L", snakefmt_bin_url, "-o", "binary.bin")
	err = snakefmt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	snakefmt_src_url := "https://files.pythonhosted.org/packages/b4/61/0228586a10b76064431e1d0c965f030b4c7dfbea6d1cfcb4d3f0cb0e6726/snakefmt-0.10.2.src.tar.gz"
	snakefmt_cmd_src := exec.Command("curl", "-L", snakefmt_src_url, "-o", "source.tar.gz")
	err = snakefmt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	snakefmt_cmd_direct := exec.Command("./binary")
	err = snakefmt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
