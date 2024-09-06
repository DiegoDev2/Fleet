package main

// Psql2csv - Run a query in psql and output the result as CSV
// Homepage: https://github.com/fphilipe/psql2csv

import (
	"fmt"
	
	"os/exec"
)

func installPsql2csv() {
	// Método 1: Descargar y extraer .tar.gz
	psql2csv_tar_url := "https://github.com/fphilipe/psql2csv/archive/refs/tags/v0.12.tar.gz"
	psql2csv_cmd_tar := exec.Command("curl", "-L", psql2csv_tar_url, "-o", "package.tar.gz")
	err := psql2csv_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	psql2csv_zip_url := "https://github.com/fphilipe/psql2csv/archive/refs/tags/v0.12.zip"
	psql2csv_cmd_zip := exec.Command("curl", "-L", psql2csv_zip_url, "-o", "package.zip")
	err = psql2csv_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	psql2csv_bin_url := "https://github.com/fphilipe/psql2csv/archive/refs/tags/v0.12.bin"
	psql2csv_cmd_bin := exec.Command("curl", "-L", psql2csv_bin_url, "-o", "binary.bin")
	err = psql2csv_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	psql2csv_src_url := "https://github.com/fphilipe/psql2csv/archive/refs/tags/v0.12.src.tar.gz"
	psql2csv_cmd_src := exec.Command("curl", "-L", psql2csv_src_url, "-o", "source.tar.gz")
	err = psql2csv_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	psql2csv_cmd_direct := exec.Command("./binary")
	err = psql2csv_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libpq")
exec.Command("latte", "install", "libpq")
}
