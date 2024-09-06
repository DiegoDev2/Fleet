package main

// Pgcli - CLI for Postgres with auto-completion and syntax highlighting
// Homepage: https://pgcli.com/

import (
	"fmt"
	
	"os/exec"
)

func installPgcli() {
	// Método 1: Descargar y extraer .tar.gz
	pgcli_tar_url := "https://files.pythonhosted.org/packages/d7/ff/10c1eb5fb8e4a81bb60abb4d13bfa04e27564fb7880915abdf603069cc93/pgcli-4.1.0.tar.gz"
	pgcli_cmd_tar := exec.Command("curl", "-L", pgcli_tar_url, "-o", "package.tar.gz")
	err := pgcli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pgcli_zip_url := "https://files.pythonhosted.org/packages/d7/ff/10c1eb5fb8e4a81bb60abb4d13bfa04e27564fb7880915abdf603069cc93/pgcli-4.1.0.zip"
	pgcli_cmd_zip := exec.Command("curl", "-L", pgcli_zip_url, "-o", "package.zip")
	err = pgcli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pgcli_bin_url := "https://files.pythonhosted.org/packages/d7/ff/10c1eb5fb8e4a81bb60abb4d13bfa04e27564fb7880915abdf603069cc93/pgcli-4.1.0.bin"
	pgcli_cmd_bin := exec.Command("curl", "-L", pgcli_bin_url, "-o", "binary.bin")
	err = pgcli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pgcli_src_url := "https://files.pythonhosted.org/packages/d7/ff/10c1eb5fb8e4a81bb60abb4d13bfa04e27564fb7880915abdf603069cc93/pgcli-4.1.0.src.tar.gz"
	pgcli_cmd_src := exec.Command("curl", "-L", pgcli_src_url, "-o", "source.tar.gz")
	err = pgcli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pgcli_cmd_direct := exec.Command("./binary")
	err = pgcli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libpq")
	exec.Command("latte", "install", "libpq").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
