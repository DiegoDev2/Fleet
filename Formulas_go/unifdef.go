package main

// Unifdef - Selectively process conditional C preprocessor directives
// Homepage: https://dotat.at/prog/unifdef/

import (
	"fmt"
	
	"os/exec"
)

func installUnifdef() {
	// Método 1: Descargar y extraer .tar.gz
	unifdef_tar_url := "https://dotat.at/prog/unifdef/unifdef-2.12.tar.gz"
	unifdef_cmd_tar := exec.Command("curl", "-L", unifdef_tar_url, "-o", "package.tar.gz")
	err := unifdef_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	unifdef_zip_url := "https://dotat.at/prog/unifdef/unifdef-2.12.zip"
	unifdef_cmd_zip := exec.Command("curl", "-L", unifdef_zip_url, "-o", "package.zip")
	err = unifdef_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	unifdef_bin_url := "https://dotat.at/prog/unifdef/unifdef-2.12.bin"
	unifdef_cmd_bin := exec.Command("curl", "-L", unifdef_bin_url, "-o", "binary.bin")
	err = unifdef_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	unifdef_src_url := "https://dotat.at/prog/unifdef/unifdef-2.12.src.tar.gz"
	unifdef_cmd_src := exec.Command("curl", "-L", unifdef_src_url, "-o", "source.tar.gz")
	err = unifdef_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	unifdef_cmd_direct := exec.Command("./binary")
	err = unifdef_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
