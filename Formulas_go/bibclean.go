package main

// Bibclean - BibTeX bibliography file pretty printer and syntax checker
// Homepage: https://www.math.utah.edu/~beebe/software/bibclean/bibclean-03.html#HDR.3

import (
	"fmt"
	
	"os/exec"
)

func installBibclean() {
	// Método 1: Descargar y extraer .tar.gz
	bibclean_tar_url := "https://ftp.math.utah.edu/pub/bibclean/bibclean-3.07.tar.xz"
	bibclean_cmd_tar := exec.Command("curl", "-L", bibclean_tar_url, "-o", "package.tar.gz")
	err := bibclean_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bibclean_zip_url := "https://ftp.math.utah.edu/pub/bibclean/bibclean-3.07.tar.xz"
	bibclean_cmd_zip := exec.Command("curl", "-L", bibclean_zip_url, "-o", "package.zip")
	err = bibclean_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bibclean_bin_url := "https://ftp.math.utah.edu/pub/bibclean/bibclean-3.07.tar.xz"
	bibclean_cmd_bin := exec.Command("curl", "-L", bibclean_bin_url, "-o", "binary.bin")
	err = bibclean_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bibclean_src_url := "https://ftp.math.utah.edu/pub/bibclean/bibclean-3.07.tar.xz"
	bibclean_cmd_src := exec.Command("curl", "-L", bibclean_src_url, "-o", "source.tar.gz")
	err = bibclean_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bibclean_cmd_direct := exec.Command("./binary")
	err = bibclean_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
