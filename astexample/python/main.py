class PersonalDetails:
    def __init__(self, name: str, surname: str):
        self.name = name
        self.surname = surname

    def __repr__(self) -> str:
        return f"PersonalDetails(name={self.name}, surname={self.surname})"


def main() -> None:
    john = PersonalDetails("John", "Doe")
    print(f"Hello, {john}!")


if __name__ == "__main__":
    main()
