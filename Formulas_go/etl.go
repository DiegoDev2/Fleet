package main

// Etl - Extensible Template Library
// Homepage: https://synfig.org

import (
	"fmt"
	
	"os/exec"
)

func installEtl() {
	// Método 1: Descargar y extraer .tar.gz
	etl_tar_url := "https://downloads.sourceforge.net/project/synfig/development/1.5.2/ETL-1.5.2.tar.gz"
	etl_cmd_tar := exec.Command("curl", "-L", etl_tar_url, "-o", "package.tar.gz")
	err := etl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	etl_zip_url := "https://downloads.sourceforge.net/project/synfig/development/1.5.2/ETL-1.5.2.zip"
	etl_cmd_zip := exec.Command("curl", "-L", etl_zip_url, "-o", "package.zip")
	err = etl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	etl_bin_url := "https://downloads.sourceforge.net/project/synfig/development/1.5.2/ETL-1.5.2.bin"
	etl_cmd_bin := exec.Command("curl", "-L", etl_bin_url, "-o", "binary.bin")
	err = etl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	etl_src_url := "https://downloads.sourceforge.net/project/synfig/development/1.5.2/ETL-1.5.2.src.tar.gz"
	etl_cmd_src := exec.Command("curl", "-L", etl_src_url, "-o", "source.tar.gz")
	err = etl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	etl_cmd_direct := exec.Command("./binary")
	err = etl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: glibmm@2.66")
exec.Command("latte", "install", "glibmm@2.66")
}
