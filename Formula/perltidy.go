package main

// Perltidy - Indents and reformats Perl scripts to make them easier to read
// Homepage: https://perltidy.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installPerltidy() {
	// Método 1: Descargar y extraer .tar.gz
	perltidy_tar_url := "https://downloads.sourceforge.net/project/perltidy/20240903/Perl-Tidy-20240903.tar.gz"
	perltidy_cmd_tar := exec.Command("curl", "-L", perltidy_tar_url, "-o", "package.tar.gz")
	err := perltidy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	perltidy_zip_url := "https://downloads.sourceforge.net/project/perltidy/20240903/Perl-Tidy-20240903.zip"
	perltidy_cmd_zip := exec.Command("curl", "-L", perltidy_zip_url, "-o", "package.zip")
	err = perltidy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	perltidy_bin_url := "https://downloads.sourceforge.net/project/perltidy/20240903/Perl-Tidy-20240903.bin"
	perltidy_cmd_bin := exec.Command("curl", "-L", perltidy_bin_url, "-o", "binary.bin")
	err = perltidy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	perltidy_src_url := "https://downloads.sourceforge.net/project/perltidy/20240903/Perl-Tidy-20240903.src.tar.gz"
	perltidy_cmd_src := exec.Command("curl", "-L", perltidy_src_url, "-o", "source.tar.gz")
	err = perltidy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	perltidy_cmd_direct := exec.Command("./binary")
	err = perltidy_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
