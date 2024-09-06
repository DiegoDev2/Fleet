package main

// Fastbit - Open-source data processing library in NoSQL spirit
// Homepage: https://sdm.lbl.gov/fastbit/

import (
	"fmt"
	
	"os/exec"
)

func installFastbit() {
	// Método 1: Descargar y extraer .tar.gz
	fastbit_tar_url := "https://web.archive.org/web/20210319090732/code.lbl.gov/frs/download.php/file/426/fastbit-2.0.3.tar.gz"
	fastbit_cmd_tar := exec.Command("curl", "-L", fastbit_tar_url, "-o", "package.tar.gz")
	err := fastbit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fastbit_zip_url := "https://web.archive.org/web/20210319090732/code.lbl.gov/frs/download.php/file/426/fastbit-2.0.3.zip"
	fastbit_cmd_zip := exec.Command("curl", "-L", fastbit_zip_url, "-o", "package.zip")
	err = fastbit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fastbit_bin_url := "https://web.archive.org/web/20210319090732/code.lbl.gov/frs/download.php/file/426/fastbit-2.0.3.bin"
	fastbit_cmd_bin := exec.Command("curl", "-L", fastbit_bin_url, "-o", "binary.bin")
	err = fastbit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fastbit_src_url := "https://web.archive.org/web/20210319090732/code.lbl.gov/frs/download.php/file/426/fastbit-2.0.3.src.tar.gz"
	fastbit_cmd_src := exec.Command("curl", "-L", fastbit_src_url, "-o", "source.tar.gz")
	err = fastbit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fastbit_cmd_direct := exec.Command("./binary")
	err = fastbit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
