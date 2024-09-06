package main

// TemporalTables - Temporal Tables PostgreSQL Extension
// Homepage: https://pgxn.org/dist/temporal_tables/

import (
	"fmt"
	
	"os/exec"
)

func installTemporalTables() {
	// Método 1: Descargar y extraer .tar.gz
	temporaltables_tar_url := "https://github.com/arkhipov/temporal_tables/archive/refs/tags/v1.2.2.tar.gz"
	temporaltables_cmd_tar := exec.Command("curl", "-L", temporaltables_tar_url, "-o", "package.tar.gz")
	err := temporaltables_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	temporaltables_zip_url := "https://github.com/arkhipov/temporal_tables/archive/refs/tags/v1.2.2.zip"
	temporaltables_cmd_zip := exec.Command("curl", "-L", temporaltables_zip_url, "-o", "package.zip")
	err = temporaltables_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	temporaltables_bin_url := "https://github.com/arkhipov/temporal_tables/archive/refs/tags/v1.2.2.bin"
	temporaltables_cmd_bin := exec.Command("curl", "-L", temporaltables_bin_url, "-o", "binary.bin")
	err = temporaltables_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	temporaltables_src_url := "https://github.com/arkhipov/temporal_tables/archive/refs/tags/v1.2.2.src.tar.gz"
	temporaltables_cmd_src := exec.Command("curl", "-L", temporaltables_src_url, "-o", "source.tar.gz")
	err = temporaltables_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	temporaltables_cmd_direct := exec.Command("./binary")
	err = temporaltables_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: postgresql@14")
	exec.Command("latte", "install", "postgresql@14").Run()
}
